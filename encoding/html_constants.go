package encoding

const resourceHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <title>UnknownResource</title>
	<style>
		body{
		    margin: 0;
    		padding: 0;
    		color: #333;
    		background-color: #eee;
    		font: 1em/1.2 'Helvetica Neue', Helvetica, Arial, Geneva, sans-serif;
		}

		h1,h2,h3 {
    		margin: 0 0 .5em;
    		font-weight: 500;
    		line-height: 1.1;
		}

		h1 { font-size: 2.25em; }
		h2 { font-size: 1.375em; }
		h3 { font-size: 1.375em; background: lightgrey; padding: 0.25em }
		
		p {
			margin: 0 0 1.5em;
			line-height: 1.5;
		}

		table {
			background-color: transparent;
			border-spacing: 0;
			border-collapse: collapse;
			border-top: 1px solid #ddd;
			width: 100%
		}
		
		th, td {
			padding: .5em 1em;
			vertical-align: center;
			text-align: left;
			border-bottom: 1px solid #ddd;
		}
		
		td:last-child {
			width:100%;
		}
		
		
		.btn {
			color: #fff !important;
			background-color: GRAY;
			border-color: #222;
			display: inline-block;
			padding: .5em 1em;
			font-weight: 400;
			line-height: 1.2;
			text-align: center;
			white-space: nowrap;
			vertical-align: middle;
			cursor: pointer;
			border: 1px solid transparent;
			border-radius: .2em;
			text-decoration: none;
		}
		
		.btn:hover {
			color: #fff !important;
			background-color: #555;
		}
		
		.btn:focus {
			color: #fff !important;
			background-color: #555;
		}
		
		.btn:active {
			color: #fff !important;
			background-color: #555;
		}
		
		.container {
			max-width: 70em;
			margin: 0 auto;
			background-color: #fff;
		}
		
		.header {
			color: #fff;
			background: #555;
			p
		}
		
		.subheader {
			color: #fff;
			background: #555;
			p
		}
		
		.header-heading { margin: 0; }
		
		.content { padding: 1em 1.25em; }
		
		.embedded { padding-left: 1.5em }
		
		@media (min-width: 42em) {
			.header { padding: 1.5em 3em; }
			.subheader { padding: .2em 3em; }
			.content { padding: 2em 3em; }
		}
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1 class="header-heading">UnknownResource</h1>
	    </div>
		<div class="content">
			<table>
				{{range $key, $value := .Values}}
				<tr>
					<td>{{$key}}</td><td>{{FormatValue $value}}</td>
				</tr>
				{{end}}
			</table>
	    </div>
    </div>
</body>
</HTML>`
