<!doctype html>
<html ng-app='rulesApp'>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

        <title>{{<.Title>}}</title>
        <link rel="Shortcut Icon" href="static/img/favicon.ico" />
        {{<range .HeadStyles>}}
            <link rel="stylesheet" href="{{<.>}}">
        {{<end>}}

        {{<range .HeadScripts>}}
            <script src="{{<.>}}"></script>
        {{<end>}}
    </head>

    <body>
      <header>
        <img class="logo" src="static/img/logoud.png"  alt="Universidad Distrital"/>
        <h1 class="app_title">{{<.Title>}}</h1>
      </header>
      <nav class="navbar navbar-inverse" role="navigation" ng-controller="parserMenu" style="width:100%;">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-ex6-collapse">
                <span class="sr-only">Desplegar navegación</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">MEN&Uacute;</a>
          </div>
          <div class="collapse navbar-collapse navbar-ex6-collapse">
            <ul class="nav navbar-nav navbar-right">
              <li class="dropdown" ng-repeat="opcion in menu">
                <a class='dropdown-toggle' data-toggle='dropdown'>
                  {{opcion.text}}<b class='caret'></b>
                </a>
                <ul class='dropdown-menu'>
                  <li ng-repeat="op in opcion.subopciones">
                    <a href="{{op.url}}" >
                      {{op.text}}
                    </a>

                  </li>
                </ul>
              </li>
            </ul>
          </div>
        </nav>
        {{<.LayoutContent>}}
    </body>
</html>
