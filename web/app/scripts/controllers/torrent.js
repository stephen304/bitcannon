'use strict';

/**
 * @ngdoc function
 * @name bitCannonApp.controller:TorrentCtrl
 * @description
 * # TorrentCtrl
 * Controller of the bitCannonApp
 */
angular.module('bitCannonApp')
  .controller('TorrentCtrl', function ($rootScope, $scope, $stateParams, $http) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $scope.btih = $stateParams.btih;
    var init = function() {
      $http.get($rootScope.api + 'torrent/' + $scope.btih).
        success(function(data, status) {
          console.log(data);
          if (status === 200) {
            $scope.torrent = data;
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
