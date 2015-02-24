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
    $scope.category = $stateParams.category;
    $scope.results = [];
    if (typeof $scope.category === 'undefined') {
      // Do nothing
    } else {
      $scope.query = $scope.query + '/c/' + $scope.category;
    }
    $scope.busy = false;
    $scope.infinite = function() {
      if ($scope.busy) {return;}
      $scope.busy = true;
      $http.get($rootScope.api + 'search/' + $scope.query + '/s/' + $scope.results.length).
        success(function(data, status) {
          if (status === 200) {
            if (data.length === 0) {
              $scope.infinite = function(){};
            }
            for (var i = 0; i < data.length; i++) {
              $scope.results.push(data[i]);
            }
            $scope.busy = false;
          }
          else {
            $scope.busy = false;
            $rootScope.message = data.message;
          }
        }).
        error(function() {
          $scope.busy = false;
          $rootScope.message = 'API Request failed.';
        });
    };
  });
