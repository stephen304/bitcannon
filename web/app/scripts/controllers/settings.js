'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:SettingsCtrl
 * @description
 * # SettingsCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('SettingsCtrl', function ($rootScope, $scope, $window) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $scope.saveAPI = function() {
      $rootScope.api = $scope.apiBox;
      $window.localStorage.api = $rootScope.api;
    };
    $scope.clearAPIBox = function() {
      $scope.apiBox = $rootScope.api;
    };
    $scope.resetAPI = function() {
      delete $window.localStorage.api;
      $rootScope.api = '';// Old default http://localhost:1337
      $scope.apiBox = $rootScope.api;
    };
    $scope.apiBox = $rootScope.api;
  });
