-- 008_add_first_last_columns.sql
-- 첫번째/마지막 번호 위치 통계 컬럼 추가

ALTER TABLE lotto_analysis_stats
ADD COLUMN IF NOT EXISTS first_count INTEGER DEFAULT 0,  -- 해당 번호가 Num1(첫번째)로 나온 누적 횟수
ADD COLUMN IF NOT EXISTS last_count INTEGER DEFAULT 0;   -- 해당 번호가 Num6(마지막)으로 나온 누적 횟수
