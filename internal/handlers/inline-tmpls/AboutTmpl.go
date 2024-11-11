package inlinetmpls

const AboutTmpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Video List</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
	<style>
		/* Custom color for the "About Me" heading */
		.about-me-title {
			color: #ff0000; /* Change to desired color */
		}
	</style>
	</head>
	<body class="d-flex flex-column min-vh-100 bg-light">

		<!-- Header -->
		<header class="bg-primary text-white py-3">
			<div class="container">
				<nav class="navbar navbar-expand-lg navbar-dark">
					<a class="navbar-brand" href="#">Video Server</a>
					<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
						<span class="navbar-toggler-icon"></span>
					</button>
					<div class="collapse navbar-collapse" id="navbarNav">
						<ul class="navbar-nav ms-auto">
							<li class="nav-item">
								<a class="nav-link active" href="/">Home</a>
							</li>
							<li class="nav-item">
								<a class="nav-link active" href="/about">About</a>
							</li>
						</ul>
					</div>
				</nav>
			</div>
		</header>


		<!-- About Application -->
		<div class="container my-3">
			<div class="card shadow-sm border-0">
				<div class="card-body">
					<h1 class="card-title text-primary mb-4">About This Application</h1>
					<p class="card-text fs-5">
						This application is a video server that allows users to stream and manage their video content.
						It is built using the Go programming language and leverages various Go libraries and frameworks 
						to provide a robust and efficient service.
					</p>
					<p class="card-text text-muted fs-6">
						<strong>Current Server IP Address:</strong> {{.IPAddress}}
					</p>
				</div>
			</div>
		</div>
	
		<!-- About Me Section -->
		<div class="container my-3">
			<div class="card shadow-sm border-0">
				<div class="card-body">
					<h1 class="card-title about-me-title mb-4">About Me</h1>
					<p class="card-text fs-5">
						Hello! I’m Rajath Kumar, a developer passionate about building efficient and robust applications.
					</p>
					<p class="fs-5">You can check out my work on <a href="https://github.com/rajath002" target="_blank">GitHub</a>.</p>
					<p class="fs-5">Feel free to reach out to me on <a href="https://www.linkedin.com/in/rajath-acharya" target="_blank">LinkedIn</a> for collaborations, questions, or just to connect!</p>
				</div>
			</div>
		</div>

		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous" defer></script>
	</body>
</html>
`
