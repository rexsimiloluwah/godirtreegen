# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Godirgen < Formula
  desc ""
  homepage "https://github.com/rexsimiloluwah/godirgen"
  version "0.1.4"

  on_macos do
    url "https://github.com/rexsimiloluwah/godirgen/releases/download/v0.1.4/godirgen_0.1.4_darwin_all.tar.gz"
    sha256 "25cecf0237159c54704c847ddaa699cc2fe60a27acd41e3cb361c118823715ea"

    def install
      bin.install "godirgen"
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/godirgen/releases/download/v0.1.4/godirgen_0.1.4_linux_arm64.tar.gz"
      sha256 "cfd012f07d0199b6f3dd98d1fdcc6d60ed42f428e698926eeef9907437ad4f6a"

      def install
        bin.install "godirgen"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/rexsimiloluwah/godirgen/releases/download/v0.1.4/godirgen_0.1.4_linux_amd64.tar.gz"
      sha256 "c3f256b5524e1ec24b758b674b86914ba82aaf77b231e00e1ac623f081a82164"

      def install
        bin.install "godirgen"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/godirgen/releases/download/v0.1.4/godirgen_0.1.4_linux_armv6.tar.gz"
      sha256 "fb284e94ff147e26881a6c98ec5e24d5884a7a36159b1dd94f359ee95924ab83"

      def install
        bin.install "godirgen"
      end
    end
  end
end
