package main

import "testing"

const testText1 = `
[[a]]
[a.b]
c = "d"
e = 1e+100
d = 1979-05-27T07:32:00Z
[[a]]
[a.b.e]
f = true
g = [{h = "i", j = "k"}, {}]
l = '''
m \n
'''
`

const parsedJson = `{"a":[{"b":{"c":"d","d":"1979-05-27T07:32:00Z","e":1e+100}},{"b":{"e":{"f":true,"g":[{"h":"i","j":"k"},{}],"l":"m \\n\n"}}}]}`

func TestToml2Json(t *testing.T) {
	json, err := tomlToJson([]byte(testText1))
	if err != nil {
		t.Log(err)
		t.Fatal("The tomlToJson function can't convert testText1.")
	}
	jsonstr := string(json)
	if jsonstr != parsedJson {
		t.Fatalf(
			`
The converted content should be the expected one below, but the actual one is not.
  The expected content:
	%s
  The actual content:
	%s
`,
			jsonstr,
			parsedJson,
		)
	}
}
