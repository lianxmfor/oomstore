pub mod error;
mod util;

use std::{collections::HashMap, path::Path};

use error::OomError;
use google::protobuf::Empty;
use oomagent::{
    oom_agent_client::OomAgentClient,
    value,
    ChannelImportRequest,
    FeatureValueMap,
    ImportRequest,
    OnlineGetRequest,
    OnlineMultiGetRequest,
    PushRequest,
    SyncRequest,
};
use tonic::{codegen::StdError, transport, Request};
use util::parse_raw_feature_values;

pub type FValue = value::Kind;

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

    pub async fn online_get_raw(
        &mut self,
        entity_key: impl Into<String>,
        features: Vec<String>,
    ) -> Result<FeatureValueMap> {
        let res = self
            .inner
            .online_get(OnlineGetRequest { entity_key: entity_key.into(), features })
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
        entity_keys: Vec<String>,
        features: Vec<String>,
    ) -> Result<HashMap<String, FeatureValueMap>> {
        let res = self
            .inner
            .online_multi_get(OnlineMultiGetRequest { entity_keys, features })
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
        group: impl Into<Option<String>>,
        description: impl Into<Option<String>>,
        revision: impl Into<Option<i64>>,
        rows: impl Iterator<Item = Vec<u8>> + Send + 'static,
    ) -> Result<u32> {
        let mut group = group.into();
        let mut description = description.into();
        let mut revision = revision.into();
        let outbound = async_stream::stream! {
            for row in rows {
                yield ChannelImportRequest{group: group.take(), description: description.take(), revision: revision.take(), row};
            }
        };
        let res = self.inner.channel_import(Request::new(outbound)).await?.into_inner();
        Ok(res.revision_id as u32)
    }

    pub async fn import(
        &mut self,
        group: impl Into<String>,
        description: impl Into<Option<String>>,
        revision: impl Into<Option<i64>>,
        input_file: impl AsRef<Path>,
        delimiter: impl Into<Option<char>>,
    ) -> Result<u32> {
        let res = self
            .inner
            .import(ImportRequest {
                group:           group.into(),
                description:     description.into(),
                revision:        revision.into(),
                input_file_path: input_file.as_ref().display().to_string(),
                delimiter:       delimiter.into().map(String::from),
            })
            .await?
            .into_inner();
        Ok(res.revision_id as u32)
    }

    pub async fn push(
        &mut self,
        entity_key: impl Into<String>,
        group: impl Into<String>,
        kv_pairs: Vec<(impl Into<String>, impl Into<value::Kind>)>,
    ) -> Result<()> {
        let mut keys = Vec::with_capacity(kv_pairs.len());
        let mut vals = Vec::with_capacity(kv_pairs.len());
        kv_pairs.into_iter().for_each(|(k, v)| {
            keys.push(k.into());
            vals.push(oomagent::Value { kind: Some(v.into()) });
        });
        self.inner
            .push(PushRequest {
                entity_key:     entity_key.into(),
                group:          group.into(),
                features:       keys,
                feature_values: vals,
            })
            .await?
            .into_inner();
        Ok(())
    }
}
