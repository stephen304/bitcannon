'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('MainCtrl', function ($scope, $state) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $scope.submit = function() {
      if ($scope.query) {
        $state.go('search', {query: $scope.query});
      }
    }
  });
