// types.ts

import { RedisClient } from 'redis';

export type CacheConfig = {
  host: string;
  port: number;
  password?: string;
  db?: number;
};

export type RedisClientConfig = {
  client?: RedisClient;
  host: string;
  port: number;
  password?: string;
  db?: number;
};

export type CacheOptions = {
  config?: CacheConfig;
  client?: RedisClient;
};

export type CacheConfigOptions = {
  host: string;
  port: number;
  password?: string;
  db?: number;
};

export type RedisClientOptions = {
  client?: RedisClient;
  host: string;
  port: number;
  password?: string;
  db?: number;
};