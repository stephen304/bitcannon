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
          if (status === 200) {
            $scope.torrent = data;
          }
        else {
          $rootScope.message = data.message;
        }
        }).
        error(function() {
          $rootScope.message = 'API Request failed.';
        });
    };
    $scope.refreshing = false;
    $scope.refresh = function() {
      if ($scope.refreshing) {
        console.log('ignored duplicate refresh request');
        return;
      }
      $scope.refreshing = true;
      $http.get($rootScope.api + 'scrape/' + $scope.btih).
        success(function(data, status) {
          if (status === 200) {
            $scope.refreshing = false;
            $scope.torrent.Swarm = data.Swarm;
            $scope.torrent.Lastmod = data.Lastmod;
          }
          else {
            $scope.refreshing = false;
            $rootScope.message = data.message;
          }
        }).
        error(function() {
          $scope.refreshing = false;
          $rootScope.message = 'API Request failed.';
        });
    };
    init();
  });
