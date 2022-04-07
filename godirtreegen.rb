# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Godirtreegen < Formula
  desc ""
  homepage "https://github.com/rexsimiloluwah/godirtreegen"
  version "0.1.5"

  on_macos do
    url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.5/godirtreegen_0.1.5_darwin_all.tar.gz"
    sha256 "353fd7e872081311180155c3deae370648796a6be536cde46d8393dbe5d56e90"

    def install
      bin.install "godirtreegen"
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.5/godirtreegen_0.1.5_linux_arm64.tar.gz"
      sha256 "1faa1841af0a3572a77d6173f8b59615f46ec4d83a565e33aa39c57ca948fb62"

      def install
        bin.install "godirtreegen"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.5/godirtreegen_0.1.5_linux_armv6.tar.gz"
      sha256 "ad99315d6e2c81089046793a5f4cb011455f6b1c02be52d0f57828196dfa4084"

      def install
        bin.install "godirtreegen"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.5/godirtreegen_0.1.5_linux_amd64.tar.gz"
      sha256 "8b3fa5a9a1ae46e529071cf9244edf30c339583984ba48eacbf57e57bac6dbb2"

      def install
        bin.install "godirtreegen"
      end
    end
  end
end
