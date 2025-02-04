{
  description = "StackState Receiver Go Client";

  nixConfig.bash-prompt = "StackState Receiver Go Client $ ";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let  
        pkgs = import nixpkgs { inherit system; overlays = [ ]; };
        pkgs-linux = import nixpkgs { system = "x86_64-linux"; overlays = [ ]; };

        # Dependencies used for both development and CI/CD
        sharedDeps = pkgs: (with pkgs; [
          bash
          go
          gotools
          diffutils # Required for golangci-lint
          golangci-lint
          openapi-generator-cli
        ]);

        darwinDevShellExtraDeps = pkgs: pkgs.lib.optionals pkgs.stdenv.isDarwin (with pkgs.darwin.apple_sdk_11_0; [
          Libsystem 
          IOKit
        ]);
      in {

        devShells = {
          dev = pkgs.mkShell {
            buildInputs = sharedDeps(pkgs) ++ darwinDevShellExtraDeps(pkgs);
          };
        };

        devShell = self.devShells."${system}".dev;

      });
}
