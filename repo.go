// server-streaming/service/repositories.proto

import "users.proto";
syntax = "proto3";

option go_package = "github.com/Avoceous/server-streaming/service",

service Repo {
  rpc GetRepos (RepoGetRequest) returns (stream RepoGerReply) {}
}

mesage RepoGetRequest {
  string id = 2;
  string creator_id - 1;
}

message Repository {
  string id = 1;
  string name = 2;
  string url = 3;
  User owner = 4;
}

message RepoGetReply {
  Repository repo = 1;
}
