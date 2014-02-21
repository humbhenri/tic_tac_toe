var ticTacToeApp = angular.module('ticTacToeApp', []);

ticTacToeApp.controller('ticTacToeCtrl', function($scope, $http) {

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

  $scope.mark = function(row, col, mark) {
      for (var i=0; i<$scope.board.length; i++) {
          for (var j=0; j<$scope.board[i].length; j++) {
              var cell = $scope.board[i][j];
              if (cell.pos[0] === row && cell.pos[1] === col) {
                  cell.mark = mark == 1 ? 'o' : 'x';
                  break;
              }
          }
      }
  }

  $scope.markCell = function(cell) {
    $http({method: 'POST',
           url: '/mark',
           data: JSON.stringify(cell),
           headers: {'Content-Type': 'application/x-www-form-urlencoded'}
          }
         )
    .success(function(data){
        console.log(data);
        for (var i=0; i<data.Pos.length; i++) {
            var pos = data.Pos[i]
            $scope.mark(pos.Row, pos.Col, pos.M);
        }
    })
    .error(function(){console.log('error')});
  }
});
