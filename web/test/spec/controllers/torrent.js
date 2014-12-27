'use strict';

describe('Controller: TorrentCtrl', function () {

  // load the controller's module
  beforeEach(module('bitCannonApp'));

  var TorrentCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    TorrentCtrl = $controller('TorrentCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
