// import z, { ZodEmail, ZodString } from "zod";

export type Author = {
    id: string;
    name: string;
    bio: string;
    age: bigint;
};

export type CreateAuthorRequest = {
    name: string;
    bio: string;
};

export type UpdateAuthorRequest = Partial<CreateAuthorRequest>;

// export const loginUserSchema = z.object({});

// var Types(
// 	int      = Type("int")
// 	i8       = Type("i8")
// 	i16      = Type("i16")
// 	i32      = Type("i32")
// 	i64      = Type("i64")
// 	uint     = Type("uint")
// 	u8       = Type("u8")
// 	u16      = Type("u16")
// 	u32      = Type("u32")
// 	u64      = Type("u64")
// 	f32      = Type("f32")
// 	f64      = Type("f64")
// 	uuid     = Type("uuid")
// 	stringt  = Type("string")
// 	bool     = Type("bool")
// 	datetime = Type("datetime")
// 	any      = Type("any")
// 	array    = Type("[]")
// )
