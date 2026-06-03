const std = @import("std");

const checkers = @import("helpers/Checkers.zig");
const helpAndVersion = @import("helpers/HelpAndVersion.zig");

pub const IoWriter = @import("helpers/IoWriter.zig").IoWriter;
const VM = @import("./core/interpreter.zig").VM;

pub fn cli(writer: IoWriter, args: []const []const u8) !u8 {
    const alloc = writer.sys.arena.allocator();

    std.debug.print("DEBUG: ARGUMENTOS RECEBIDOS:\n", .{});
    for (args) |a| {
        try writer.stdout.interface.print("  {s}\n", .{a});
        try writer.stdout.flush();
    }

    if (args.len == 0) {
        try helpAndVersion.help(writer);
        return 1;
    }

    const arg = args[0];
    {
        // -h | --help
        if (checkers.flagsEqual(arg, &.{ "-h", "--help" })) {
            helpAndVersion.help(writer) catch {
                return 1;
            };
            return 0;
        }

        // -v | --version
        if (checkers.flagsEqual(arg, &.{ "-v", "--version" })) {
            helpAndVersion.version(writer) catch {
                return 1;
            };
            return 0;
        }
    }

    const file_path = arg;

    std.Io.Dir.cwd().access(writer.sys.io, file_path, .{}) catch |err| {
        switch (err) {
            error.FileNotFound => {
                std.debug.print("Não Existe\n", .{});
                return 1;
            },
            else => {
                std.debug.print("Erro no entry: {any}", .{err});
                return 2;
            },
        }
    };

    const content = try std.Io.Dir.cwd().readFileAlloc(writer.sys.io, file_path, alloc, std.Io.Limit.limited(1024 * 1024));
    defer alloc.free(content);

    std.debug.print("Lido: {d} bytes\n", .{content.len});
    std.debug.print("\n----\n\n", .{});

    var vm = VM.init(content, file_path, writer);
    try vm.entry(alloc);

    return 0;
}
