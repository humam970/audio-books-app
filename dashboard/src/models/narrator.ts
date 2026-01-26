export type Narrator = {
    id: string;
    name: string;
    bio: string | null;
};

export type CreateNarratorRequest = {
    name: string;
    bio?: string | null;
};

export type UpdateNarratorRequest = Partial<CreateNarratorRequest>;
