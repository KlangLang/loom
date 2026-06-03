const std = @import("std");
const loom = @import("loom");
const IoWriter = loom.IoWriter;

pub fn main(init: std.process.Init) !void {
    var out_buff: [1024]u8 = undefined;
    var err_buff: [1024]u8 = undefined;

    var stdout_writer = std.Io.File.stdout().writer(init.io, &out_buff);
    var stderr_writer = std.Io.File.stderr().writer(init.io, &err_buff);

    var writer: IoWriter = .{
        .sys = init,

        .stdout = &stdout_writer,
        .stderr = &stderr_writer,
    };
    defer {
        writer.stdout.flush() catch {};
        writer.stderr.flush() catch {};
    }

    const args = try init.minimal.args.toSlice(init.arena.allocator());

    if (try loom.cli(writer, args[1..]) != 0) {
        try writer.stderr.interface.print("Erro: ", .{});
        return;
    }
}
