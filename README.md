<h1>gRPC-server-side-streaming</h1>
<p>Initiates Protobuf specification for the Repo service<p>
<p>Compiled by @avoceous<p> 

<img src="./resources/source/images/avoceous.png" width="500">


<o1>
// server-streaming/service/repositories.proto

import "users.proto";
syntax = "proto3";

option go_package = "github.com/username/server-streaming/service",

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
</o1>
