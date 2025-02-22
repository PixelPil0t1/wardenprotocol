export interface Env {
  EVMOS_NODE_RPC: string;
  EVMOS_REGISTRY_ADDRESS: string;
  EVMOS_PUBLIC_CLIENT_TIMEOUT_MSEC: number;
  ETHEREUM_NODE_RPC: string;
  WARDEN_EVM_CHAIN_ID: number;
  WARDEN_SIGN_REQUESTS_POLLING_INTERVAL_MSEC: number;
  WARDEN_SIGN_REQUESTS_PAGE_SIZE: number;
  SIGN_REQUESTS_PROCESSOR_SEEN_CACHE_ELEMENTS_SIZE: number;
}