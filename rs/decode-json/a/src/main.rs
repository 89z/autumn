use tinyjson::JsonValue;

let s = r#"
    {
        "bool": true,
        "arr": [1, null, "test"],
        "nested": {
            "blah": false,
            "blahblah": 3.14
        },
        "unicode": "\u2764"
    }
"#;

let parsed: JsonValue = s.parse().unwrap();
println!("Parsed: {:?}", parsed);
