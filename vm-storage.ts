import { getConfig, setEnvPaths } from './config'

setEnvPaths(['.env.test.vm.go', '.go'])

beforeEach(() => {
  getConfig({ reload: true })
})
