<!DOCTYPE html>
<html lang="zh-CN">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
  <meta name="generator" content="Jekyll v4.0.1">
  <meta name="keywords" content="{{ .keywords }}" />
  <meta name="description" content="{{ .description }}" />
  <title>{{.title}}</title>

  <!-- Bootstrap core CSS -->
  <!-- CSS only -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.0/dist/css/bootstrap.min.css" integrity="sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk" crossorigin="anonymous">
  <link href="https://cdn.bootcss.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
  <link href="https://cdn.bootcss.com/highlight.js/9.15.8/styles/solarized-dark.min.css" rel="stylesheet">
  <!--
      <link href="/assets/css/basic.css" rel="stylesheet">
      <link href="/assets/css/style.css" rel="stylesheet">
   -->
</head>

<body class="bg-light">
  <div class="container page-top animated fadeInDown">
    <div>
        <div class="content">
            {{ template "content" . }}
            <div class="pagination">
                <ul class="clearfix">
                </ul>
            </div>
        </div>
    </div>

    <footer class="my-5 pt-5 text-muted text-center text-small">
      <p class="mb-1">&copy; 2019-2020 SDUT Computer College Software Lab</p>
      <ul class="list-inline">
        <li class="list-inline-item"><a href="#">Contact</a></li>
        <li class="list-inline-item"><a href="#">Github</a></li>
        <li class="list-inline-item"><a href="#">Support</a></li>
      </ul>
    </footer>
  </div>
<script src="https://cdn.jsdelivr.net/npm/jquery@3.5.1/dist/jquery.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.0/dist/js/bootstrap.min.js" integrity="sha384-OgVRvuATP1z7JjHLkuOU7Xw704+h835Lr+6QL9UvYjZE3Ipu6Tp75j7Bh/kR0JKI" crossorigin="anonymous"></script>
<script src="https://cdn.bootcss.com/highlight.js/9.15.8/highlight.min.js"></script>
<script src="https://cdn.bootcss.com/highlight.js/9.15.8/languages/go.min.js"></script>

<script type="text/javascript">
    $(document).ready(function() {
        $('pre code').each(function(i, block) {
            hljs.highlightBlock(block);
        });
    });
</script>

</body>
</html>
