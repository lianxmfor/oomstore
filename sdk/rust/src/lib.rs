pub mod error;
mod util;

use std::collections::HashMap;

use error::OomError;
use google::protobuf::Empty;
use oomagent::{
    oom_agent_client::OomAgentClient,
    value,
    ChannelImportRequest,
    FeatureValueMap,
    OnlineGetRequest,
    OnlineMultiGetRequest,
    SyncRequest,
};
use tonic::{codegen::StdError, transport, Request};
use util::parse_raw_feature_values;

type Result<T> = std::result::Result<T, OomError>;

pub mod google {
    pub mod protobuf {
        tonic::include_proto!("google.protobuf");
    }
}

pub mod oomagent {
    tonic::include_proto!("oomagent");
}

pub struct Client {
    inner: OomAgentClient<transport::Channel>,
}

impl Client {
    pub async fn connect<D>(dst: D) -> Result<Self>
    where
        D: std::convert::TryInto<tonic::transport::Endpoint>,
        D::Error: Into<StdError>,
    {
        Ok(Self { inner: OomAgentClient::connect(dst).await? })
    }

    pub async fn health_check(&mut self) -> Result<()> {
        Ok(self.inner.health_check(Empty {}).await.map(|_| ())?)
    }

    pub async fn online_get_raw(&mut self, key: impl Into<String>, features: Vec<String>) -> Result<FeatureValueMap> {
        let res = self
            .inner
            .online_get(OnlineGetRequest { entity_key: key.into(), feature_full_names: features })
            .await?
            .into_inner();
        Ok(match res.result {
            Some(res) => res,
            None => FeatureValueMap::default(),
        })
    }

    pub async fn online_get(
        &mut self,
        key: impl Into<String>,
        features: Vec<String>,
    ) -> Result<HashMap<String, value::Kind>> {
        let rs = self.online_get_raw(key, features).await?;
        Ok(parse_raw_feature_values(rs))
    }

    pub async fn online_multi_get_raw(
        &mut self,
        keys: Vec<String>,
        features: Vec<String>,
    ) -> Result<HashMap<String, FeatureValueMap>> {
        let res = self
            .inner
            .online_multi_get(OnlineMultiGetRequest { entity_keys: keys, feature_full_names: features })
            .await?
            .into_inner();
        Ok(res.result)
    }

    pub async fn online_multi_get(
        &mut self,
        keys: Vec<String>,
        features: Vec<String>,
    ) -> Result<HashMap<String, HashMap<String, value::Kind>>> {
        let rs = self.online_multi_get_raw(keys, features).await?;
        Ok(rs.into_iter().map(|(k, v)| (k, parse_raw_feature_values(v))).collect())
    }

    pub async fn sync(&mut self, revision_id: u32, purge_delay: u32) -> Result<()> {
        self.inner
            .sync(SyncRequest {
                revision_id: i32::try_from(revision_id)?,
                purge_delay: i32::try_from(purge_delay)?,
            })
            .await?;
        Ok(())
    }

    pub async fn channel_import(
        &mut self,
        group_name: impl Into<Option<String>>,
        description: impl Into<Option<String>>,
        revision: impl Into<Option<i64>>,
        rows: impl Iterator<Item = Vec<u8>> + Send + 'static,
    ) -> Result<u32> {
        let mut group_name = group_name.into();
        let mut description = description.into();
        let mut revision = revision.into();
        let outbound = async_stream::stream! {
            for row in rows {
                yield ChannelImportRequest{group_name: group_name.take(), description: description.take(), revision: revision.take(), row};
            }
        };
        let res = self.inner.channel_import(Request::new(outbound)).await?.into_inner();
        Ok(res.revision_id as u32)
    }
}
