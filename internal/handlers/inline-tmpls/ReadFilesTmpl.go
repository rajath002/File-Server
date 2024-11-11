package inlinetmpls

const ReadFilesTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Video List</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
<style>
	.f-system-ui {
		font-family: system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
	}
	.video-list-title {
		font-weight: 100;
		color: #000000;
		font-size: 1.3rem;
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
                            <a class="nav-link active" href="#">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link active" href="/about">About</a>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>
    </header>

    <!-- Main Content -->
    <main class="flex-grow-1">
        <div class="container py-5">
            <h1 class="display-4 text-center text-primary mb-5">Video List</h1>
            <ul class="list-group">
				{{ $index := 0 }}
                {{ range .VideoFiles }}
				{{ $index = add $index 1}}
                <li data-fileName="{{.}}" class="list-group-item d-flex justify-content-between align-items-center">
                    <span>
						<span class="fw-bold fs-5">{{$index}}.</span>
						&nbsp; <span class="video-list-title">{{.}}</span>
						</span>
					<div>
						<a target="_blank" href="/video-player/{{.}}" class="btn btn-outline-primary btn-sm">Play</a>
						<a href="/videos/{{.}}" download class="btn btn-outline-primary btn-sm">Download</a>
					</div>
                </li>
                {{end}}
            </ul>
        </div>
    </main>

    <!-- Footer -->
    <footer class="bg-dark text-white py-3">
        <div class="container d-flex justify-content-center">
            <span>&copy; 2024 Video Server. All Rights Reserved.</span>
        </div>
    </footer>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous" defer></script>
	<script>
	document.addEventListener('DOMContentLoaded', function() {
		document.querySelectorAll('li[data-fileName]').forEach(function(element) {
			element.addEventListener('click', handleFileDownload);
		});
	});

	window.addEventListener('beforeunload', function() {
		console.log('before unload')
		document.querySelectorAll('li[data-fileName]').forEach(function(element) {
			element.removeEventListener('click', handleFileDownload);
		});
	});

	function handleFileDownload(event) {
		const fileName = this.getAttribute('data-fileName');
		// Add your file download logic here
		console.log('Clicked filename : ', fileName);
		this.classList.add('border-success')
	}
	</script>
</body>
</html>`
