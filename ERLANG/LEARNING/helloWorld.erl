-module(helloWorld).
% -export([start/0, while/1]).
-export([main/0]).

% start()->
%     A = 12,
%     B = 2.3e-3,
%     C = atom1,
%     D = "Test for string", 
%     io:fwrite("Hello World, A: ~w, B: ~w, C: ~w, D: ~p~n", [A, B, C, D]).


while(L) -> 
    io:fwrite("~w", [L]).

main() ->
    while([1, 2, 3]).