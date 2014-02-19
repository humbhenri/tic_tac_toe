var ticTacToeApp = angular.module('ticTacToeApp', []);

ticTacToeApp.controller('ticTacToeCtrl', function($scope, $http) {

  $scope.player1 = 'x';
  $scope.player2 = 'o';
  $scope.player = $scope.player1;

  $scope.board = [
    [
      {'mark' : '', 'pos' : [0, 0]},
      {'mark' : '', 'pos' : [0, 1]},
      {'mark' : '', 'pos' : [0, 2]}
    ],
    [
      {'mark' : '', 'pos' : [1, 0]},
      {'mark' : '', 'pos' : [1, 1]},
      {'mark' : '', 'pos' : [1, 2]}
    ],
    [
      {'mark' : '', 'pos' : [2, 0]},
      {'mark' : '', 'pos' : [2, 1]},
      {'mark' : '', 'pos' : [2, 2]}
    ]

  ];

  $scope.markCell = function(cell) {
    $http({method: 'POST',
           url: '/mark',
           data: JSON.stringify(cell),
           headers: {'Content-Type': 'application/x-www-form-urlencoded'}
          }
         )
    .success(function(data){console.log(data)})
    .error(function(){console.log('error')});
  }
});