{{ define "content" }}

<div class="archive animated fadeInDown">
    <div class="list-with-title">
        {{ range .tags }}
        <div class="listing-title">
            <a href="{{ .Url }}">
                {{ .Name }}({{ .Count }})
            </a>

        </div>
        {{ end }}
    </div>
</div>

{{ end }}