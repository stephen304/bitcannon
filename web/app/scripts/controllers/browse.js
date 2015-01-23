'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:BrowseCtrl
 * @description
 * # BrowseCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('BrowseCtrl', function ($rootScope, $scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
