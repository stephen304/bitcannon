'use strict';

describe('Controller: BrowsesearchCtrl', function () {

  // load the controller's module
  beforeEach(module('bitCannonApp'));

  var BrowsesearchCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    BrowsesearchCtrl = $controller('BrowsesearchCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
