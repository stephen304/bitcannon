'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:SearchCtrl
 * @description
 * # SearchCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('SearchCtrl', function($rootScope, $scope, $stateParams, $http) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $scope.query = $stateParams.query;
    var init = function() {
      $http.get($rootScope.api + '/search/' + $scope.query).
        success(function(data, status) {
          console.log(data);
          if (status === 200) {
            $scope.results = data;
          }
        else {
          // Error!
        }
        }).
        error(function() {
          // Error!
        });
    };
    init();
  });
