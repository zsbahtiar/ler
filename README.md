# ler

ler syntax for teler custom rules.

## Specifications

Specification for ler syntax:

### Rule Definition:

 - A rule is defined using the following syntax:
   ```
   rule "<rule name>":
     condition: <condition>
     [
       { method: <method>, element: <element>, pattern: "<pattern>" }
       ...
     ];
   ```
 - The rule name is enclosed in quotes (single/double).
 - The condition specifies the logical condition for the rule and accepts only two values: `$OR` or `$AND`.
 - The rules are enclosed in square brackets (`[]`), and each rule is defined as a key-value pair enclosed in curly braces (`{}`).
 - Each rule must have the following properties:
   - `method` (or `m`): Specifies the HTTP method for the rule and accepts the following values: `$GET`, `$HEAD`, `$POST`, `$PUT`, `$PATCH`, `$DELETE`, `$CONNECT`, `$OPTIONS`, `$TRACE`, `$ALL`. (refer to [request.Method](https://pkg.go.dev/github.com/kitabisa/teler-waf/request#Method) type)
   - `element` (or `el`, `e`): Specifies the target element for the rule and accepts the following values: `$URI`, `$HEADERS`, `$BODY`, `$ANY`. (refer to [`request.Element`](https://pkg.go.dev/github.com/kitabisa/teler-waf/request#Element) type)
   - `pattern` (or `p`): Specifies the regular expression pattern to match.
 - Each rule definition must be terminated with a semicolon (`;`).

#### Notes

 - The rule name is a descriptive name for the rule and should be unique.
 - The condition specifies how the rules within a group should be evaluated. If `$OR` is used, any of the rules within the group being evaluated can be true for the overall condition to be true. If `$AND` is used, all the rules within the group must be true for the overall condition to be true.
 - The method property specifies the HTTP method for which the rule should apply. It can be one of the predefined values: `$GET`, `$HEAD`, `$POST`, `$PUT`, `$PATCH`, `$DELETE`, `$CONNECT`, `$OPTIONS`, `$TRACE`, or `$ALL`. The `$ALL` value means the rule applies to all HTTP methods. (refer to [`request.ALL`](https://pkg.go.dev/github.com/kitabisa/teler-waf/request#ALL) constant)
 - The element property specifies the target element for the rule. It can be one of the predefined values: `$URI`, `$HEADERS`, `$BODY`, or `$ANY`. The `$ANY` value means the rule applies to any element. (refer to [`request.Any`](https://pkg.go.dev/github.com/kitabisa/teler-waf/request#Any) constant)
 - The pattern property specifies a regular expression pattern that the target element should match for the rule to be true enclosed in quotes.

### Syntax Rules

 - The rule definition starts with the keyword `rule` (or `r`) followed by the rule name enclosed in quotes.
 - The condition is specified using the keyword `condition` (or `cond`, `c`) followed by the condition value (`$OR` or `$AND`).
 - Each rule has properties specified as key-value pairs: `method: <method>`, `element: <element>`, and `pattern: "<pattern>"`. The properties are separated by commas (`,`).
 - The method, element, and pattern values are case-sensitive.

### Examples:
 - Example 1:
   ```
   rule "Rule name":
     condition: $OR
     [
       { method: $GET, element: $URI, pattern: "\$\{.*://.*\/?\w+?}" }
       { method: $GET, element: $HEADERS, pattern: "curl" }
     ];
   ```
 - Example 2:
   ```
   rule "New rule":
     condition: $AND
     [
       { method: $POST, element: $URI, pattern: "/api/posts" }
       { method: $ANY, element: $HEADERS, pattern: "authorization" }
       { method: $GET, element: $BODY, pattern: "^true|false$" }
     ];
   ```
 - Example 3:
   ```
   r "Rule 3":
     c: $AND
     [
       { m: $ANY, e: $URI, p: "/\.bash" }
     ];
   ```