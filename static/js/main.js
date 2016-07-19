// Creación del módulo
var rulesApp = angular.module('rulesApp', ['datatables']);

rulesApp.controller('homeController', function($scope) {
  $scope.message = 'Versión 0.0.1';
});

rulesApp.controller('dominioController', function($scope, $http) {
  $scope.title = 'Dominios';
  $scope.message = 'Listado de Dominios de Aplicación';

  $http.get('http://localhost:8081/v1/dominio?limit=0')
  .then(function(response) {
      $scope.data = response.data;
  });
});
rulesApp.controller('agregarDominioController', function($scope, $http) {
  $scope.title = 'Dominios';
  $scope.message = 'Agregar Dominio de Aplicación';

  $scope.add = function(){
    if($scope.domain.name == null){
      return;
    }
    if($scope.domain.description == null){
      return;
    }
    var data = {
        Nombre: $scope.domain.name,
        Descripcion: $scope.domain.description
    };
    $http.post('http://localhost:8081/v1/dominio',data);
  };
});

rulesApp.controller('predicadoController', function($scope, $http) {
  $scope.title = 'Predicados';
  $scope.message = 'Listado de Predicados de Aplicación';

  $http.get('http://localhost:8081/v1/predicado?limit=0')
  .then(function(response) {
      $scope.data = response.data;
  });
});
rulesApp.controller('agregarPredicadoController', function($scope, $http) {
  $scope.title = 'Predicados';
  $scope.message = 'Agregar Predicados de Aplicación';

  $http.get('http://localhost:8081/v1/dominio?limit=0')
  .then(function(response) {
      $scope.dominios = response.data;
  });

  $http.get('http://localhost:8081/v1/tipo_predicado/?limit=0')
  .then(function(response) {
      $scope.tipos = response.data;
  });

  $scope.add = function(){
    if($scope.predicado.name == null){
      return;
    }
    if($scope.predicado.description == null){
      return;
    }
    var data = {
        Dominio: {Id: $scope.predicado.dominio.id},
        Nombre: $scope.predicado.name,
        Descripcion: $scope.predicado.description,
        TipoPredicado: {Id: $scope.predicado.tipopredicado.id}
    };
    $http.post('http://localhost:8081/v1/predicado',data);
  };
});

rulesApp.controller('aboutController', function($scope) {
  $scope.message = 'Esta es la página "Acerca de"';
});

rulesApp.controller('parserMenu', function($scope, $http) {
  $scope.menu = [{"id": "1", "text": "Dominios", "url": "",
                  "subopciones": [{ "text": "Listar", "url": "/listarDominios" },
                                  { "text": "Agregar", "url": "/agregarDominios" }]
              },
              {"id": "2", "text": "Predicados", "url": "",
                  "subopciones": [{ "text": "Listar", "url": "/listarPredicados" },
                                  { "text": "Agregar", "url": "/agregarPredicados" }]
              }];

});
