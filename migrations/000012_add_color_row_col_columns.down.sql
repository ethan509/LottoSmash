-- 012_add_color_row_col_columns.sql (down)
-- 색상/행/열 통계 컬럼 삭제

ALTER TABLE lotto_analysis_stats
DROP COLUMN IF EXISTS color_count,
DROP COLUMN IF EXISTS color_prob,
DROP COLUMN IF EXISTS row_count,
DROP COLUMN IF EXISTS row_prob,
DROP COLUMN IF EXISTS col_count,
DROP COLUMN IF EXISTS col_prob;
