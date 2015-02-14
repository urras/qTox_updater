using Go = import "go.capnp";
$Go.package("main");

@0xa934be1c402e8e56;

struct ValidateFile {
	name @0: Text;
	hash @1: Data;
	buildnum @2: UInt8;
	sign @3: Data;
}

struct FileControl {
	name @0: Text;
	hashnext @1: Data;
	chunknum @2: UInt8;
	chunksize @3: UInt32 = 1024;
	sign @4: Data;
}

struct FileChunk {
	chunknum @0: UInt8;
	data @1: Data;
	hashnext @2: Data;
}

struct GetManifest {
	buildnum @0: UInt8;
}

struct SendManifest {
	buildnum @0: UInt8;
	data @1: Data;
	sign @2: Data;
}