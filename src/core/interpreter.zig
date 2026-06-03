const std = @import("std");
const Opcodes = @import("opcodes.zig").Opcodes;

pub const IoWriter = @import("../helpers/IoWriter.zig").IoWriter;

fn Stack(comptime T: type) type {
    return struct {
        list: std.ArrayList(T),

        const Self = @This();

        pub fn init() Self {
            return .{
                .list = .empty,
            };
        }

        pub fn push(self: *Self, alloc: std.mem.Allocator, value: T) !void {
            try self.list.append(alloc, value);
        }

        pub fn pop(self: *Self) !T {
            return self.list.pop() orelse error.StackUnderFlow;
        }

        pub fn deinit(self: *Self, alloc: std.mem.Allocator) !void {
            self.list.deinit(alloc);
        }
    };
}

pub const VM = struct {
    global_stack: Stack(u32),
    code: []const u8,
    file_path: []const u8,
    idx: usize = 0,

    writer: IoWriter,

    pub fn init(code: []const u8, file_path: []const u8, writer: IoWriter) VM {
        return .{
            .global_stack = Stack(u32).init(),
            .code = code,
            .file_path = file_path,
            .writer = writer,
        };
    }

    pub fn entry(self: *VM, alloc: std.mem.Allocator) !void {
        // validacao besta so pra ver
        if (self.code.len < 0) {
            std.debug.print(".fiber invalido\n", .{});
            return;
        }

        while (true) {
            // fetch
            const op = Opcodes.toOp(self.readNextByte());

            switch (op) {
                .PUSH_I8 => {
                    const n1 = self.readNextByte();
                    try self.global_stack.push(alloc, n1);

                    continue;
                },

                .PUSH_I16 => {
                    const n1 = self.readNextByte();
                    try self.global_stack.push(alloc, n1);

                    continue;
                },

                .PUSH_I32 => {
                    const n1 = self.readNextByte();
                    try self.global_stack.push(alloc, n1);

                    continue;
                },

                .PUSH_I64 => {
                    const n1 = self.readNextByte();
                    try self.global_stack.push(alloc, n1);

                    continue;
                },

                .PUSH_F16 => {
                    const n1 = self.readNextByte();
                    try self.global_stack.push(alloc, n1);
                    continue;
                },

                .PUSH_F32 => {
                    const n1 = self.readNextByte();
                    try self.global_stack.push(alloc, n1);

                    continue;
                },

                .PUSH_F64 => {
                    const n1 = self.readNextByte();

                    try self.global_stack.push(alloc, n1);

                    continue;
                },

                // adds
                .ADD_I8 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                .ADD_I16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                .ADD_I32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                .ADD_I64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                .ADD_F16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                .ADD_F32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                .ADD_F64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 + n2);
                    continue;
                },

                // subs
                .SUB_I8 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                .SUB_I16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                .SUB_I32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                .SUB_I64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                .SUB_F16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                .SUB_F32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                .SUB_F64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 - n2);
                    continue;
                },

                // muls
                .MUL_I8 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                .MUL_I16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                .MUL_I32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                .MUL_I64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                .MUL_F16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                .MUL_F32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                .MUL_F64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    try self.global_stack.push(alloc, n1 * n2);
                    continue;
                },

                // divs
                .DIV_I8 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .DIV_I16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .DIV_I32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .DIV_I64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .DIV_F16 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .DIV_F32 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .DIV_F64 => {
                    const n1 = try self.global_stack.pop();
                    const n2 = try self.global_stack.pop();

                    if (n2 == 0) {
                        try self.panic("The denominator do not equal 0.");
                        break;
                    }

                    try self.global_stack.push(alloc, n1 / n2);
                    continue;
                },

                .CALL_NATIVE => {
                    const built_in_code = self.readNextByte();

                    if (built_in_code == 0x00) {
                        try self.writer.stdout.interface.print("{any}\n", .{try self.global_stack.pop()});
                    }

                    continue;
                },

                .HALT => break,
                else => break,
            }
        }
    }

    fn panic(self: VM, msg: []const u8) !void {
        try self.writer.stderr.interface.print("LOOM PANIC: {s}", .{msg});
    }

    fn readNextByte(self: *VM) u8 {
        if (self.idx >= self.code.len) return 0;

        const byte = self.code[self.idx];
        self.idx += 1;

        return byte;
    }
};
