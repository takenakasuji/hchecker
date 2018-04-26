# hchecker
hcheker is bodyless browser.
This tool is useful if you want to easily check the http header.

## Demo
### Output type inline
![demo](https://raw.githubusercontent.com/takenakasuji/hchecker/master/pictures/inline_format_example.gif)

### Output type json
![demo](https://raw.githubusercontent.com/takenakasuji/hchecker/master/pictures/json_format_example.gif)

## Command Line Options
### -url
Enter URL to request
Note: This option is mandatory

```
./hchecker -url https://google.com
``` 

### -format
Select display format type

#### Usage
```
# Display line type
# Inline type is the default, so if you do'nt select the inline option, it will automatically be of this type.
./hchecker -url https://google.com -formt inline

# Display json type
./hchecker -url https://google.com -formt json
``` 

## Features

