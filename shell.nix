let
  pkgs = import <nixpkgs> { };
in
pkgs.stdenv.mkDerivation {
  name = "rmapi";
  buildInputs = with pkgs; [
    go
    gotools
  ];
}
