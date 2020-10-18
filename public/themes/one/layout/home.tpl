{{ define "content" }}
<div class="py-5 text-center">
  <img class="d-block mx-auto mb-4" src="/assets/bootstrap-solid.svg" alt="" width="72" height="72">
  <h2>Checkout form</h2>
  <p class="lead">Below is an example form built entirely with Bootstrap’s form controls. Each required form group has a validation state that can be triggered by attempting to submit the form without completing it.</p>
</div>

<div class="row">
  <div class="col-md-12 order-md-12 mb-12">
    <table class="table table-bordered">
      <thead>
        <tr>
          <th scope="col" colspan="5">Title</th>
          <th scope="col" colspan="1" class="text-center">Date</th>
          <th scope="col" colspan="1" class="text-center">Author</th>
        </tr>
      </thead>
      <tbody>
        {{ range .artlist }}
            <tr>
              <td colspan="5">
                <a href="{{ .Url }}">{{ .Title }}</a>
              </td>
              <td colspan="1" class="text-center">{{ .CreatedAt | format }}</td>
              <td colspan="1" class="text-center">{{ .Author }}</td>
            </tr>
        {{ end }}
      </tbody>
    </table>
  </div>
</div>
<div class="row">
  <div class="col justify-content-md-start">

  </div>
  <div class="col justify-content-md-end" style="text-align: right;">
    <button type="button" class="btn btn-outline-info">Prev</button>
    <button type="button" class="btn btn-outline-info">Next</button>
  </div>
</div>
<div class="pagination">
    <ul class="clearfix">
    </ul>
</div>
{{ end }}