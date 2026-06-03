const std = @import("std");

pub const Opcodes = enum(u8) {
    NOP = 0x00,

    // push`s
    PUSH_I8 = 0x10,
    PUSH_I16 = 0x11,
    PUSH_I32 = 0x12,
    PUSH_I64 = 0x13,

    PUSH_F16 = 0x14,
    PUSH_F32 = 0x15,
    PUSH_F64 = 0x16,

    // add`s
    ADD_I8 = 0x20,
    ADD_I16 = 0x21,
    ADD_I32 = 0x22,
    ADD_I64 = 0x23,

    ADD_F16 = 0x24,
    ADD_F32 = 0x25,
    ADD_F64 = 0x26,

    //sub`s
    SUB_I8 = 0x30,
    SUB_I16 = 0x31,
    SUB_I32 = 0x32,
    SUB_I64 = 0x33,

    SUB_F16 = 0x34,
    SUB_F32 = 0x35,
    SUB_F64 = 0x36,

    // mul`s
    MUL_I8 = 0x40,
    MUL_I16 = 0x41,
    MUL_I32 = 0x42,
    MUL_I64 = 0x43,

    MUL_F16 = 0x44,
    MUL_F32 = 0x45,
    MUL_F64 = 0x46,

    // div`s
    DIV_I8 = 0x50,
    DIV_I16 = 0x51,
    DIV_I32 = 0x52,
    DIV_I64 = 0x53,

    DIV_F16 = 0x54,
    DIV_F32 = 0x55,
    DIV_F64 = 0x56,

    // fn op`s
    CALL_NATIVE = 0x70,
    CALL = 0x71,
    RET = 0x72,

    HALT = 0xFF,

    pub fn toOp(bin: u8) Opcodes {
        inline for (@typeInfo(Opcodes).@"enum".fields) |field| {
            if (field.value == bin) return @enumFromInt(bin);
        }
        return .NOP;
    }
};
