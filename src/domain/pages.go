package domain

const IndexPages = `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<title>{{.Title}}</title>

			<!-- Bootstrap CSS -->
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-0evHe/X+R7YkIZDRvuzKMRqM+OrBnVFBL6DOitfPri4tjfHxaWutUpFmBp4vmVor" crossorigin="anonymous">
		</head>
		<body class="p-3">
			<h1>{{.Title}}</h1>

			<!-- 
				. in top context is data (variable passed when executing the template)
				.Tables -> data.Tables
			-->
			{{range .Tables}}
				<h2>{{.TableCategory}}</h2>
				<table class="table table-bordered mt-3">
					<thead>
						<tr>
							<th scope="col">X</th>
							<!-- 
								. here means data.Tables[index]
								.TableHeader -> data.Tables[index].TableHeader
								{{.}} Access each element in TableHeader slice
							-->
							{{range .TableHeader}}
								<th scope="col">{{.}}</th>
							{{end}}
						</tr>
					</thead>
					<tbody>
							<!-- 
								. here means data.Tables[index]
								.TableRow -> data.Tables[index].TableRow
								{{.}} Access each element in TableRow slice
							-->
						{{range .TableRow}}
							<tr>
								<!-- 
									. here means data.Tables[index].TableRow[index]
									{{.}} Access each element in TableRow[index] slice
								-->
								{{range .}}
									<td>{{.}}</td>
								{{end}}
							</tr>
						{{end}}
					</tbody>
				</table>
			{{end}}

			<!-- Bootstrap JS -->
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js" integrity="sha384-pprn3073KE6tl6bjs2QrFaJGz5/SUsLqktiwsUTF55Jfv3qYSDhgCecCxMW52nD2" crossorigin="anonymous"></script>
		</body>
	</html>
`
