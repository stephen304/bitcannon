'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:BrowseCtrl
 * @description
 * # BrowseCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('BrowseCtrl', function ($rootScope, $scope, $http) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var init = function() {
      $http.get($rootScope.api + 'browse').
        success(function(data, status) {
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
