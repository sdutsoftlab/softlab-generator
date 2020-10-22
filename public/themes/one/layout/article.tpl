{{ define "content" }}
<link href="/assets/css/next-helvetica.css" rel="stylesheet">
<div class="post-page">
    <div class="post animated fadeInDown">
        <div class="post-title">
            <h2>{{ .article.Title }}</h2>
        </div>
        <div class="progress" style="height: 5px;">
          <div class="progress-bar" role="progressbar" style="width: 100%;" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100"></div>
        </div>
        <div class="post-content" id="content">

            {{ .article.Content | unescaped }}
            <div class="card">
              <div class="card-header">
                <i class="fa fa-calendar"></i>
                {{ .article.CreatedAt | format }}
                &nbsp;&nbsp;
                <i class="fa fa-folder-open"></i>
                  {{ range .article.Category }}
                <a href="/category/{{ . }}">{{ . }}</a>&nbsp;
                {{ end }}
                &nbsp;&nbsp;
                <i class="fa fa-tags"></i>
                {{ range .article.Tags }}
                <a href="/tag/{{ . }}">{{ . }}</a>&nbsp;
                {{ end }}
              </div>
              <div class="card-body">
                <blockquote class="blockquote mb-0">
                  <p>当你的才华撑不起你的野心时，就静下心来学习吧.</p>
                  <footer class="blockquote-footer"><cite title="Source Title">网络</cite></footer>
                </blockquote>
              </div>
            </div>
        </div>

    </div>

    <!--评论-->

</div>

{{ end }}