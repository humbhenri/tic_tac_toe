var ticTacToeApp = angular.module('ticTacToeApp', []);

ticTacToeApp.controller('ticTacToeCtrl', function($scope, $http) {
    $scope.gameMsg = function(msg) {
        document.getElementById('gameMsg').innerHTML = msg;
    }
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

    $scope.markBoard =  function(pos) {
        for (var i=0; i<pos.length; i++) {
            $scope.mark(pos[i].Row, pos[i].Col, pos[i].M);
        }
    };


    $scope.markCell = function(cell) {
        $http({method: 'POST',
            url: '/mark',
            data: JSON.stringify(cell),
            headers: {'Content-Type': 'application/x-www-form-urlencoded'}
        })
        .success(function(data){
            $scope.gameMsg('Player O');
            console.log(data);
            
            if (data.Pos !== undefined) {
                $scope.markBoard(data.Pos);
            } else if (data.Winner !== undefined) {
                $scope.markBoard(data.Board.Pos);
                if (data.Winner == 0) {
                    $scope.gameMsg('Game Over: Draw !');
                    console.log('draw');
                } else {
                    var winner = data.Winner == 1 ? 'o' : 'x';
                    $scope.gameMsg('Game Over: ' + winner + ' wins !');
                }
            } else {
                $scope.gameMsg('Player X');
            }
            })
        .error(function(data){
            console.log(data);
        });
    }
});
