-- 012_add_color_row_col_columns.sql
-- 색상/행/열 통계 컬럼 추가

ALTER TABLE lotto_analysis_stats
ADD COLUMN IF NOT EXISTS color_count INTEGER DEFAULT 0,        -- 해당 번호 색상의 총 출현 횟수
ADD COLUMN IF NOT EXISTS color_prob DOUBLE PRECISION DEFAULT 0, -- 색상 출현 확률
ADD COLUMN IF NOT EXISTS row_count INTEGER DEFAULT 0,          -- 해당 번호 행의 총 출현 횟수
ADD COLUMN IF NOT EXISTS row_prob DOUBLE PRECISION DEFAULT 0,   -- 행 출현 확률
ADD COLUMN IF NOT EXISTS col_count INTEGER DEFAULT 0,          -- 해당 번호 열의 총 출현 횟수
ADD COLUMN IF NOT EXISTS col_prob DOUBLE PRECISION DEFAULT 0;   -- 열 출현 확률
