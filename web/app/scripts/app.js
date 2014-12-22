'use strict';

/**
 * @ngdoc overview
 * @name bitCannonApp
 * @description
 * # bitCannonApp
 *
 * Main module of the application.
 */
angular
  .module('bitCannonApp', [
    'ngAnimate',
    'ui.router'
  ])
    .config(function ($stateProvider, $urlRouterProvider) {
      $urlRouterProvider.otherwise('/');
      $stateProvider
        .state('index', {
          url: '/',
          templateUrl: 'views/main.html',
          controller:'MainCtrl',
          pageTitle: 'Home'
        });
      });
