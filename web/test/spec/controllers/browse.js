'use strict';

describe('Controller: BrowseCtrl', function () {

  // load the controller's module
  beforeEach(module('bitCannonApp'));

  var BrowseCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    BrowseCtrl = $controller('BrowseCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
