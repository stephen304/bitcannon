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
    'ui.router',
    'angular-loading-bar'
  ])
  .config(function($stateProvider, $urlRouterProvider, $compileProvider) {
    $compileProvider.aHrefSanitizationWhitelist(/^\s*(https?|ftp|mailto|chrome-extension|magnet):/);
    $urlRouterProvider.otherwise('/');
    $stateProvider
      .state('index', {
        url: '/',
        templateUrl: 'views/main.html',
        controller: 'MainCtrl',
        pageTitle: 'Home'
      })
      .state('browse', {
        url: '/browse',
        templateUrl: 'views/browse.html',
        controller: 'BrowseCtrl',
        pageTitle: 'Browse'
      })
      .state('search', {
        url: '/search/:query',
        templateUrl: 'views/search.html',
        controller: 'SearchCtrl',
        pageTitle: 'Search'
      })
      .state('torrent', {
        url: '/torrent/:btih',
        templateUrl: 'views/torrent.html',
        controller: 'TorrentCtrl',
        pageTitle: 'Torrent'
      })
      .state('about', {
        url: '/about',
        templateUrl: 'views/about.html',
        controller: 'MainCtrl',
        pageTitle: 'About'
      });
  })
  .run(function($rootScope) {
    $rootScope.api = 'http://localhost:1337';
  });
