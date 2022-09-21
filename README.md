<h1>gRPC-server-side-streaming</h1>
<p>Initiates Protobuf specification for the Repo service<p>
<p>Compiled by @avoceous<p> 

<img src="./resources/source/images/avoceous.png" width="500">





<o1>
<li>// server-streaming/service/repositories.proto

<li>import "users.proto";
<li>syntax = "proto3";

<li>option go_package = "github.com/username/server-streaming/service",

<li>service Repo {
  rpc GetRepos (RepoGetRequest) returns (stream RepoGerReply) {}
}

<li>mesage RepoGetRequest {
  string id = 2;
  string creator_id - 1;
}

<li>message Repository {
  string id = 1;
  string name = 2;
  string url = 3;
  User owner = 4;
}

<li>message RepoGetReply {
  Repository repo = 1;
}
</o1>
