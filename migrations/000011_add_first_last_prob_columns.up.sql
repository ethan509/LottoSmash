-- 011_add_first_last_prob_columns.sql
-- 첫번째/마지막 번호 위치 확률 컬럼 추가

ALTER TABLE lotto_analysis_stats
ADD COLUMN IF NOT EXISTS first_prob DOUBLE PRECISION DEFAULT 0,  -- 해당 번호가 Num1(첫번째)로 나올 확률
ADD COLUMN IF NOT EXISTS last_prob DOUBLE PRECISION DEFAULT 0;   -- 해당 번호가 Num6(마지막)으로 나올 확률
