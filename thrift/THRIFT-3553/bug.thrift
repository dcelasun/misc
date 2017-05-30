namespace go bug

struct TestsResponse {
	1: map<string, Test> tests
}

struct Test {
	1: string id
	2: string name
}
