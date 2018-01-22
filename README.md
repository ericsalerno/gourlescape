# gourlescape

A super simple golang URL escape library (for practice). 

## Usage

By default the URL escaper will use a hyphen "-" character to replace characters, lower case to replace upper case, and leave the trailing replacements in the string. Example usage for URL escaping would be:

    e := Escape{}
	value := e.URL(" Some URL escaP3d text!!  	")
    fmt.Println(value) //some-url-escap3d-text--

You can also change that behavior as follows:

    e := Escape{
        Replacer: "_",
        RemoveTail: true,
    }
    value := e.URL(" Some URL escaP3d text!!  	")
    fmt.Println(value) //some_url_escap3d_text