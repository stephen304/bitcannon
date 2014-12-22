'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('MainCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
