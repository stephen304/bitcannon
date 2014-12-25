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
        })
        .state('search', {
          url: '/search',
          templateUrl: 'views/search.html',
          controller:'SearchCtrl',
          pageTitle: 'Search'
        })
        .state('about', {
          url: '/about',
          templateUrl: 'views/about.html',
          controller:'MainCtrl',
          pageTitle: 'About'
        });
      });
