# num2text
Golang library to convert number to text in different ways

Use it if you want to convert number to text, ordinal, or full ordinal.

<h4>Numbers to Text</h4>

```go
num2text.ConvertToText(1)    // one
num2text.ConvertToText(2)    // two
num2text.ConvertToText(999)  // nine hundred ninety nine
```

<h4>Numbers to Ordinal</h4>

```go
num2text.ConvertToOrdinal(1)    // 1st
num2text.ConvertToOrdinal(2)    // 2nd
num2text.ConvertToOrdinal(999)  // 999th
```

<h4>Numbers to Ordinal with Full description</h4>

```go
num2text.ConvertToOrdinalFull(1)    // first
num2text.ConvertToOrdinalFull(2)    // second
num2text.ConvertToOrdinalFull(999)  // nine hundred ninety nine thousand nine hundred ninety nineth
num2text.ConvertToOrdinalFull(55)   // fifty fifth
num2text.ConvertToOrdinalFull(105)  // one hundred fifth
```

<h4>Text to Number convertion</h4>

```go
num2text.ConvertToNumber("ten thousand twelfth") // 10012
num2text.ConvertToNumber("ten thousand twelve") // 10012
num2text.ConvertToNumber("10012th") // 10012
```
