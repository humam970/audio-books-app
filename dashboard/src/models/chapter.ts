export type Chapter = {
    id: string;
    book_id: string;
    title: string;
    start_time: number;
    end_time: number;
    order_index: number;
};

export type CreateChapterRequest = {
    title: string;
    start_time: number;
    end_time: number;
    order_index: number;
};

export type UpdateChapterRequest = Partial<CreateChapterRequest>;
