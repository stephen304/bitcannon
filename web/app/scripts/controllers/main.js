'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('MainCtrl', function($scope, $state) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $scope.submit = function() {
      if ($scope.query) {
        if ($scope.selectedCategory) {
          $state.go('searchCategory', {
            query: $scope.query,
            category: $scope.selectedCategory.name
          });
        } else {
          $state.go('search', {
            query: $scope.query
          });
        }
      }
    };
  });
